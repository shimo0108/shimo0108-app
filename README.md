# shimo0108-app
＊使用技術
バックエンド Go(Echo)
フロントエンド Vue.js(vuex,vuetify)
インフラ構築 terraform
インフラ AWS ECS(fargate) ALB CloudWatch ECR S3 CloudFront RDS(postgres)

＊日々の予定を書き込めるカレンダーアプリです。
https://shimo0108-app.com

<img width="1783" alt="スクリーンショット 2021-11-11 22 35 55" src="https://user-images.githubusercontent.com/60634601/141307433-244528d8-7030-43f7-8678-ecaddff87958.png">

アピールポイント
バックエンド
・ORM使わず実装(現在実務ではゴリゴリORMを使用しているRailsを使っているため、ORMなしでもDBに関する基本的な知識は持っている事をアピールしたかった)
・DB依存をなくしたテスト環境を構築するためgo-sqlmockを採用(テスト時に本物のDBを使う必要がなくなるため、「DBに入っているデータをバックアップ→テストに必要な前提データを挿入→テスト終わったらバックアップしたデータを元に戻す」みたいなことが必要なくなる。そのため、DBを使った関数のUT(IT)を手軽に高速にテストできる。すごくおすすめ)

フロントエンド
・Vuexを用いたstate管理、lint/formatterを用いたコード自動整形
・UI/UXを意識したデザイン

インフラ
・ECS(fargate)を使用
・terraformでインフラ環境完全コード化

その他
・SPA実装(Vue.js & Go)
・github actionsを用いたCI/CD


https://github.com/shimo0108/task_list
こちらは上記ポートフォリオとほぼ同内容のアプリですが、フロントからバックエンド間の通信をgRPC-Web、プロキシにenvoyを採用しています。
https://www.envoyproxy.io/docs/envoy/latest/intro/what_is_envoy

課題
・突貫工事で作成したのでコード汚い、リファクタリングする。
・テストコードがまだ少ない。DB周りやapi周りのテストも充実させる。
・https://shimo0108-app.comにログインユーザー用のGoサーバーをもう一つ増築予定。現在のserverコンテナとgRPC通信で通信を行えるよう改良する。プロキシとしてenvoyを用いる。(App-Mesh使いたい)
