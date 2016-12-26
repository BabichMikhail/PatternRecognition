<!DOCTYPE html>
<html>
<head>
</head>
<body>
    <form method="GET">
        <div>
            <label>Размерность</label></br>
            <input name="D" value="{{ .D }}"/>
        </div>
        <div>
            <label>Число кластеров</label></br>
            <input name="N" value="{{ .N }}"/>
        </div>
        <div>
            <label>Параметр T</label></br>
            <input name="T" value="{{ .T }}"/>
        </div>
        <div>
            <label>Параметр K</label></br>
            <input name="K" value="{{ .K }}"/>
        </div>
        <!--<div>
            <label>Суммарное число точек</label></br>
            <input name="NN" value="{{ .NN }}"/>
        </div>-->
        <div>
            <label>Параметр Alpha</label></br>
            <input name="Alpha" value="{{ .Alpha }}"/>
        </div>
        <div>
            <label>Компоненты вектора</label></br>
            <input name="Comp1" value="{{ .Comp1 }}"/>
            <input name="Comp2" value="{{ .Comp2 }}"/>
        <div>
            <button type="submit" name="update" value="1">Обновить</button>
        </div>
        <div>
            <table>
                <thead>
                    <th></th>
                    <th>Норм. корреляционная матрица</th>
                    <th colspan="{{ .D }}">Вектор дисперсий</th>
                    <th colspan="{{ .D }}">Вектор средних</th>
                    <th>Число точек</th>
                    <!--<th>Априорная вероятность</th>-->
                </thead>
                <tbody>
                    {{ range $key, $cluster := .Clusters }}
                    <tr>
                        <td>Кластер #{{ $key }}</td>
                        <td>
                            <table>
                                <tbody>
                                    {{ range $i, $corRow := $cluster.Cor }}
                                    <tr>
                                        {{ range $j, $cor := $corRow }}
                                        <td><input name="cluster-сor-{{ $key }}-{{ $i }}-{{ $j }}" value="{{ $cor }}"/></td>
                                        {{ end }}
                                    </tr>
                                    {{ end }}
                                </tbody>
                            </table>
                        </td>
                        {{ range $i, $disp := $cluster.Disp }}
                        <td><input name="cluster-disp-{{ $key }}-{{ $i }}" value="{{ $disp }}"/></td>
                        {{ end }}
                        {{ range $i, $med := $cluster.Med }}
                        <td><input name="cluster-med-{{ $key }}-{{ $i }}" value="{{ $med }}"/></td>
                        {{ end }}
                        <td><input name="cluster-N-{{ $key }}" value="{{ $cluster.N }}"/></td>
                        <!--<td><input name="cluster-P-{{ $key }}" value="{{ $cluster.P }}"/></td>-->
                    </tr>
                    {{ end }}
                </tbody>
            </table>
        </div>
        <!--<div>
            <label>Границы:</label>
            <div>X: <input name="left_x" value="{{ .LeftX }}"/><input name="right_x" value="{{ .RightX }}"/></div>
            <div>Y: <input name="bottom_y" value="{{ .BottomY }}"/><input name="top_y" value="{{ .TopY }}"/></div>
        </div>-->
        <input type="submit" value="Применить"/>
    </form>
    {{ range $key, $value := .Data0 }}
    <div>
        <p>Исходные данные ({{ len $value }} класт.)</p>
        <table>
            <tr>
                <td>
                    <img src="/static/img/img_lab8_0.png" width="512">
                </td>
                <td>
                    <table border="1">
                        <thead>
                            <th>Количество точек</th>
                            <th colspan="{{ $.D }}">Средние</th>
                            <th colspan="{{ $.D }}">Дисперсии</th>
                        </thead>
                        <tbody>
                            {{ range $key2, $value2 := $value }}
                            <tr>
                                <td>{{ $value2.N }}</td>
                                {{ range $idx, $med := $value2.Med }}
                                <td>{{ $med }}</td>
                                {{ end }}
                                {{ range $idx, $disp := $value2.Disp }}
                                <td>{{ $disp }}</td>
                                {{ end }}
                            </tr>
                            {{ end }}
                        </tbody>
                    </table>
                </td>
            </tr>
        </table>
    </div>
    {{ end }}
    {{ range $key, $value := .Data1 }}
    <div>
        <p>T - алгоритм ({{ len $value }} класт.)</p>
        <table>
            <tr>
                <td>
                    <img src="/static/img/img_lab8_1.png" width="512">
                </td>
                <td>
                    <table border="1">
                        <thead>
                            <th>Количество точек</th>
                            <th colspan="{{ $.D }}">Средние</th>
                            <th colspan="{{ $.D }}">Дисперсии</th>
                        </thead>
                        <tbody>
                            {{ range $key2, $value2 := $value }}
                            <tr>
                                <td>{{ $value2.N }}</td>
                                {{ range $idx, $med := $value2.Med }}
                                <td>{{ $med }}</td>
                                {{ end }}
                                {{ range $idx, $disp := $value2.Disp }}
                                <td>{{ $disp }}</td>
                                {{ end }}
                            </tr>
                            {{ end }}
                        </tbody>
                    </table>
                </td>
            </tr>
        </table>
    </div>
    {{ end }}
    {{ range $key, $value := .Data2 }}
    <div>
        <p>Алгоритм максиминного расстояния ({{ len $value }} класт.)</p>
        <table>
            <tr>
                <td>
                    <img src="/static/img/img_lab8_2.png" width="512">
                </td>
                <td>
                    <table border="1">
                        <thead>
                            <th>Количество точек</th>
                            <th colspan="{{ $.D }}">Средние</th>
                            <th colspan="{{ $.D }}">Дисперсии</th>
                        </thead>
                        <tbody>
                            {{ range $key2, $value2 := $value }}
                            <tr>
                                <td>{{ $value2.N }}</td>
                                {{ range $idx, $med := $value2.Med }}
                                <td>{{ $med }}</td>
                                {{ end }}
                                {{ range $idx, $disp := $value2.Disp }}
                                <td>{{ $disp }}</td>
                                {{ end }}
                            </tr>
                            {{ end }}
                        </tbody>
                    </table>
                </td>
            </tr>
        </table>
    </div>
    {{ end }}
    <div>
        {{ $idx := index .Data3 0 }}
        <p>Алгоримт K-внутригрупповых средних  ({{ len $idx }} класт.) ({{ len .Data3 }} итерац.)</p>
        {{ range $key, $value := .Data3 }}
        <table>
            <tr>
                <td>
                    <img src="/static/img/img_lab8_3_{{ $key }}.png" width="512">
                </td>
                <td>
                    <table border="1">
                        <thead>
                            <th>Количество точек</th>
                            <th colspan="{{ $.D }}">Средние</th>
                            <th colspan="{{ $.D }}">Дисперсии</th>
                        </thead>
                        <tbody>
                            {{ range $key2, $value2 := $value }}
                            <tr>
                                <td>{{ $value2.N }}</td>
                                {{ range $idx, $med := $value2.Med }}
                                <td>{{ $med }}</td>
                                {{ end }}
                                {{ range $idx, $disp := $value2.Disp }}
                                <td>{{ $disp }}</td>
                                {{ end }}
                            </tr>
                            {{ end }}
                        </tbody>
                    </table>
                </td>
            </tr>
        </table>
        {{ end }}
    </div>
</body>
</html>
