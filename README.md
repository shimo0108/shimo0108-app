# shimo0108-app
＊使用技術<br>
バックエンド Go(Echo)<br>
フロントエンド Vue.js(vuex,vuetify)<br>
インフラ構築 terraform<br>
インフラ AWS ECS(fargate) ALB CloudWatch ECR S3 CloudFront RDS(postgres)<br>

＊日々の予定を書き込めるカレンダーアプリです。<br>
https://shimo0108-app.com<br>

<img width="1783" alt="スクリーンショット 2021-11-11 22 35 55" src="https://user-images.githubusercontent.com/60634601/141307433-244528d8-7030-43f7-8678-ecaddff87958.png">

*インフラ構成図<br>
![Untitled Diagram drawio](https://user-images.githubusercontent.com/60634601/144709465-a183248b-6a90-41cd-b946-3e6d0bcae873.png)<br>


アピールポイント<br>
バックエンド<br>
・ORM使わず実装(現在実務ではゴリゴリORMを使用しているRailsを使っているため、ORMなしでもDBに関する基本的な知識は持っている事をアピールしたかった)<br>
・DB依存をなくしたテスト環境を構築するためgo-sqlmockを採用(テスト時に本物のDBを使う必要がなくなるため、「DBに入っているデータをバックアップ→テストに必要な前提データを挿入→テスト終わったらバックアップしたデータを元に戻す」みたいなことが必要なくなる。そのため、DBを使った関数のUT(IT)を手軽に高速にテストできる。すごくおすすめ)<br>

フロントエンド<br>
・Vuexを用いたstate管理、lint/formatterを用いたコード自動整形<br>
・UI/UXを意識したデザイン<br>

インフラ<br>
・ECS(fargate)を使用<br>
・terraformでインフラ環境完全コード化<br>

その他<br>
・SPA実装(Vue.js & Go)<br>
・github actionsを用いたCI/CD<br>


https://github.com/shimo0108/task_list<br>
こちらは上記ポートフォリオとほぼ同内容のアプリですが、フロントからバックエンド間の通信をgRPC-Web、プロキシにenvoyを採用しています。<br>

envoyについては下記サイトを参考<br>
https://www.envoyproxy.io/docs/envoy/latest/intro/what_is_envoy<br>

課題<br>
・突貫工事で作成したのでコード汚い、リファクタリングする。<br>
・テストコードがまだ少ない。DB周りやapi周りのテストも充実させる。<br>
・ログインユーザー用のGoサーバーをもう一つ増築予定。現在のserverコンテナとgRPC通信で通信を行えるよう改良する。
・プロキシとしてenvoyを用いる。(App-Mesh使いたい)
