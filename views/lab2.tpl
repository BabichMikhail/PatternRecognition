<!DOCTYPE html>

<html>
<head>
  <title>Lab2.2</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <style type="text/css">

  </style>
</head>

<body>
    <header>
    </header>
    <content>
        <form method="GET">
            <div>
                <label for="N">Размер выборки (N):</label>
                <input id="N" value="{{ .N }}" name="N"/>
            </div>
            <div>
                <label for="gist">Количество интервалов гистограммы:</label>
                <input id="gist" value="{{ .GistIntervals }}" name="Gist"/>
            </div>
            <button type="submit">Обновить</button>
        </form>
        <table>
            <thead>
                <th></th>
                <th>Эксперимент</th>
                <th>Анализ</th>
                <th>Разница</th>
            </thead>
            <tbody>
                <tr align='center'>
                    <td>Дисперсия</td>
                    <td>{{ .Disp }}</td>
                    <td>{{ .DispAn }}</td>
                    <td>{{ .DispDif }}</td>
                </tr>
                <tr align='center'>
                    <td>Среднее</td>
                    <td>{{ .Srednee }}</td>
                    <td>{{ .SredneeAn }}</td>
                    <td>{{ .SredneeDif }}</td>
                </tr>
                <tr align='center'>
                    <td>Коэф. ассиметрии</td>
                    <td>{{ .Assim }}</td>
                    <td>{{ .AssimAn }}</td>
                    <td>{{ .AssimDif }}</td>
                </tr>
                <tr align='center'>
                    <td>Коэф. эксцесса</td>
                    <td>{{ .Ekscess }}</td>
                    <td>{{ .EkscessAn }}</td>
                    <td>{{ .EkscessDif }}</td>
                </tr>
          </tbody>
      </table>
      <canvas height='1600' width='1570px' id='canvas' style="border:1px solid #000000;"></canvas>
  </content>
  <footer>
  </footer>
</body>
<script type="text/javascript">
    var canvas = document.getElementById("canvas")
	var	ctx = canvas.getContext('2d')
    var N = {{ .N }}
    var intervals = {{ .Intervals }}
    var a = []
    {{ range $key, $value := .Gist }}
    a[{{$key}}] = {{ $value }}
    {{ end }}
    var L = {{ .L }}
    var width = 40;
    if ({{ .Intervals}} > 25) {
        canvas.width = 20 + width*(intervals - 1) + 36 + 50;

    }
    ctx.fillStyle = "#000000"
    var k = 1400.0 / N;
    ctx.font= "12px Arial"
    for (let i = 0; i < intervals; ++i) {
        let height = k * a[i];
        ctx.fillRect(20 + width*(i - 1) + 38, 480 - height, 28, height)
        ctx.fillText(a[i], 20 + width*(i - 1) + 36, 480 - height - 10)
        ctx.fillText(Math.round((L*i + L/2)*1000)/1000, 20 + width*(i - 1) + 36, 480 + 20)
    }
    ctx.beginPath();
    ctx.moveTo(20 - width + 28 + 6, 480)
    ctx.lineTo(20 - width + 28 + canvas.width - 20, 480)
    ctx.lineTo(20 - width + 28 + canvas.width - 20 - 10, 480 - 4)
    ctx.moveTo(20 - width + 28 + canvas.width - 20, 480)
    ctx.lineTo(20 - width + 28 + canvas.width - 20 - 10, 480 + 4)
    ctx.moveTo(20 - width + 28 + 6, 480)
    ctx.lineTo(20 - width + 28 + 6, 100)
    ctx.lineTo(20 - width + 28 + 6 + 4, 100 + 10)
    ctx.moveTo(20 - width + 28 + 6, 100)
    ctx.lineTo(20 - width + 28 + 6 - 4, 100 + 10)
    ctx.stroke()
    ctx.fillText({{ .A }}, 20 - width + 28 + 6 - 5, 480 + 11);
    ctx.fillText({{ .B }}, 20 + width*(intervals - 1) + 36 - 10 , 480 + 11);

    ctx.beginPath();
    var lastX = 20 - width + 28 + 6
    var lastY = 1080 - 4
    ctx.moveTo(lastX, lastY)
    var prev = 0;
    for (let i = 0; i < intervals; ++i) {
        prev += a[i]
        let height = k * a[i];
        ctx.lineTo(20 + width*(i - 1) + 34, 1080 - height)
        lastX = 20 + width*(i + 0) + 34
        lastY = 1080 - height
        ctx.fillText(Math.round(prev / N * 1000) / 1000, 20 + width*(i - 1) + 36, 1080 - height - 10)
        ctx.fillText(Math.round((L*i + L/2)*1000)/1000, 20 + width*(i - 1) + 36, 1080 + 20)
        ctx.lineTo(lastX, lastY)

    }
    ctx.lineTo(lastX + 100, lastY)
    ctx.stroke()

    ctx.beginPath()
    ctx.moveTo(20 - width + 28 + 6, 1080)
    ctx.lineTo(20 - width + 28 + canvas.width - 20, 1080)
    ctx.lineTo(20 - width + 28 + canvas.width - 20 - 10, 1080 - 4)
    ctx.moveTo(20 - width + 28 + canvas.width - 20, 1080)
    ctx.lineTo(20 - width + 28 + canvas.width - 20 - 10, 1080 + 4)
    ctx.moveTo(20 - width + 28 + 6, 1080)
    ctx.lineTo(20 - width + 28 + 6, 700)
    ctx.lineTo(20 - width + 28 + 6 + 4, 700 + 10)
    ctx.moveTo(20 - width + 28 + 6, 700)
    ctx.lineTo(20 - width + 28 + 6 - 4, 700 + 10)
    ctx.stroke()
    ctx.fillText({{ .A }}, 20 - width + 28 + 6 - 5, 1080 + 11);
    ctx.fillText({{ .B }}, 20 + width*(intervals - 1) + 36 - 10 , 1080 + 11);

</script>
</html>
