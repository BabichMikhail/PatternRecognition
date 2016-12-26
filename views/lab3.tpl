<!DOCTYPE html>
<html>
<head>
</head>
<body>
    <form method="GET">
        <div>
            <label>Размерность</label></br>
            <input name="M" value="{{ .M }}"/>
            <button type="submit" name="updateM" value="1">Обновить размерность</button>
        </div>
        <div>
            <label>Средние</label></br>
            {{ range $key, $value := .Med }}
            <input name="med{{ $key }}" value="{{ $value }}"/>
            {{ end }}
        </div>
        <div>
            <label>Корреляционная матрица</label>
            {{ range $key, $value := .Cor }}
                <div>
                {{ range $key2, $value2 := .}}
                <input name="cor{{ $key }}{{ $key2 }}" value="{{ $value2 }}"/>
                {{ end }}
                </div>
            {{ end }}
        </div>
        <div>
            <label>Компоненты случайного вектора</label></br>
            <input name="comp1" value="{{ .Comp1 }}"/>
            <input name="comp2" value="{{ .Comp2 }}"/>
        </div>
        <!--<div>
            <label>Выбрать 2 случайные компоненты</label></br>
            <input type="checkbox" name="rand" value="{{ .Rand }}"/>
        </div>-->
        <div>
            <label>Число точек</label></br>
            <input name="N-points" value="{{ .NPoints }}"/>
        </div>
        <div>
            <label>Границы:</label>
            <div>X: <input name="left_x" value="{{ .LeftX }}"/><input name="right_x" value="{{ .RightX }}"/></div>
            <div>Y: <input name="left_y" value="{{ .BottomY }}"/><input name="right_y" value="{{ .TopY }}"/></div>
        </div>
        <input type="submit" value="Применить"/>
    </form>
    <div>
        <label>Нормированная матрица</label></br>
        <table>
        {{ range $key, $value := .NormMatrix }}
            <tr>
            {{ range $key2, $value2 := $value }}
                <td>{{ $value2 }}</td>
            {{ end }}
            </tr>
        {{ end }}
        </table>
        <label>Дисперсии</label></br>
        <table>
            <tr>
            {{ range $key, $value := .Disp }}
            <td>{{ $value }}</td>
            {{ end }}
            </tr>
        </table>
    </div>
    <div>
        <img src="/static/img/img1.png">
        <img src="/static/img/img2.png">
        <img src="/static/img/img3.png">
    </div>
</body>
</html>
