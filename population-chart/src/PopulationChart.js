import React, { useEffect, useRef } from 'react';
import * as d3 from 'd3';
import Papa from 'papaparse';
import populationData from './data.csv';

const PopulationChart = () => {
  const chartRef = useRef();

  useEffect(() => {
    // Parse the CSV data
    Papa.parse(populationData, {
      header: true,
      download: true,
      complete: (result) => {
        const data = result.data;
        createChart(data);
      },
    });
  }, []);

  const createChart = (data) => {
    // Specify the chart’s dimensions (except for the height).
    const width = 928;
    const marginTop = 30;
    const marginRight = 10;
    const marginBottom = 0;
    const marginLeft = 30;

    // Determine the series that need to be stacked.
    const series = d3.stack()
      .keys(d3.union(data.map(d => d.age))) // distinct series keys, in input order
      .value(([, D], key) => D.get(key).population) // get value for each series key and stack
      (d3.index(data, d => d.state, d => d.age)); // group by stack then series key

    // Compute the height from the number of stacks.
    const height = series[0].length * 25 + marginTop + marginBottom;

    // Prepare the scales for positional and color encodings.
    const x = d3.scaleLinear()
      .domain([0, d3.max(series, d => d3.max(d, d => d[1]))])
      .range([marginLeft, width - marginRight]);

    const y = d3.scaleBand()
      .domain(d3.groupSort(data, D => -d3.sum(D, d => d.population), d => d.state))
      .range([marginTop, height - marginBottom])
      .padding(0.08);

    const color = d3.scaleOrdinal()
      .domain(series.map(d => d.key))
      .range(d3.schemeSpectral[series.length])
      .unknown("#ccc");

    // A function to format the value in the tooltip.
    const formatValue = x => isNaN(x) ? "N/A" : x.toLocaleString("en")

    // Create the SVG container.
    const svg = d3.create("svg")
      .attr("width", width)
      .attr("height", height)
      .attr("viewBox", [0, 0, width, height])
      .attr("style", "max-width: 100%; height: auto;");

    // Append a group for each series, and a rect for each element in the series.
    svg.append("g")
      .selectAll()
      .data(series)
      .join("g")
      .attr("fill", d => color(d.key))
      .selectAll("rect")
      .data(D => D.map(d => (d.key = D.key, d)))
      .join("rect")
      .attr("x", d => x(d[0]))
      .attr("y", d => y(d.data[0]))
      .attr("height", y.bandwidth())
      .attr("width", d => x(d[1]) - x(d[0]))
      .append("title")
      .text(d => `${d.data[0]} ${d.key}\n${formatValue(d.data[1].get(d.key).population)}`);

    // Append the horizontal axis.
    svg.append("g")
      .attr("transform", `translate(0,${marginTop})`)
      .call(d3.axisTop(x).ticks(width / 100, "s"))
      .call(g => g.selectAll(".domain").remove());

    // Append the vertical axis.
    svg.append("g")
      .attr("transform", `translate(${marginLeft},0)`)
      .call(d3.axisLeft(y).tickSizeOuter(0))
      .call(g => g.selectAll(".domain").remove());

    // Return the chart with the color scale as a property (for the legend).
    return Object.assign(svg.node(), { scales: { color } });
  }
};

export default PopulationChart;
