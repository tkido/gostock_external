package page

var graphTmpl = `<h3>%s</h3>
<div id="chart_%s"></div>`

var jsTmpl = `// Create the data table.
var data = new google.visualization.DataTable();
data.addColumn('string', 'Tpic');
data.addColumn('number', 'Slices');
data.addRows([%s]);

// Instantiate and draw our chart, passing in some options.
var chart = new google.visualization.PieChart(document.getElementById('chart_%s'));
chart.draw(data, options);`

var pageTmpl = `<html lang="jp">
<head>
  <meta charset="utf-8">
  <meta name="robots" content="noindex,nofollow">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{ID}} {{名称}}{{date}}</title>
  <!-- Bootstrap core CSS -->
  <link href="css/bootstrap.min.css" rel="stylesheet">
  <!-- Custom styles for this template -->
  <link href="css/stock.css" rel="stylesheet">
</head>
<body>
  <div class="container">
	<div class="starter-template">
	  <h2>{{ID}} {{名称}}{{date}}</h2>
	  <p>{{特色}}</p>
	  <!-- Div that will hold the pie chart -->
	  {{divs}}
	  <div class="charts">
		<h4>日足6ヶ月</h4>
		<img src="https://chart.yahoo.co.jp/?code={{ID}}.{{市記号}}&amp;tm=6m&amp;type=c&amp;log=off&amp;size=n&amp;over=s,v,m25&amp;add=&amp;comp=" title="日足6か月 25d平均線">
		<h4>週足2年</h4>
		<img src="https://chart.yahoo.co.jp/?code={{ID}}.{{市記号}}&amp;tm=2y&amp;type=c&amp;log=off&amp;size=n&amp;over=s,v,m75&amp;add=&amp;comp=" title="週足2年 75d平均線">
		<h4>月足10年</h4>
		<img src="https://chart.yahoo.co.jp/?code={{ID}}.{{市記号}}&amp;tm=ay&amp;type=c&amp;log=off&amp;size=n&amp;over=s,v,m260&amp;add=&amp;comp=" title="月足10年 52w平均線">
	  </div>

		<img src="https://shares.gmo-click.com/valuation/graphs/bs_{{ID}}_10.png" />
		<img src="https://shares.gmo-click.com/valuation/graphs/plh_{{ID}}_10.png" />
		<img src="https://shares.gmo-click.com/valuation/graphs/cf_{{ID}}_10.png" />

	  <!-- Load the AJAX API -->
	  <script type="text/javascript" src="https://www.google.com/jsapi"></script>
	  <script type="text/javascript">
		// Load the Visualization API and the piechart package.
		google.load('visualization', '1.0', {'packages':['corechart']});

		// Set a callback to run when the Google Visualization API is loaded.
		google.setOnLoadCallback(drawChart);

		// Callback that creates and populates a data table,
		// instantiates the pie chart, passes in the data and
		// draws it.
		function drawChart() {
		  // Set chart options
		  var options = {'width':800,
						 'height':400};
		  {{jss}}
		}
	  </script>
	</div>
  </div>
</body>
</html>
`
