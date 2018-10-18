## 抽象化のレベル

クラウド基盤を作成するにあたっての前提として、我々がやりたいことはサービスの提供・デプロイを自動化することである。よって、これまでサービスを提供するにあたって行ってきたことを綺麗にモデリングすれば自ずと正しく抽象化されると考えた。

### サービスを提供するために必要な作業

#### 十分に仮想化された場合 (VM)

1. リソースを確保 (VMやDockerで)
2. 依存パッケージのインストール (OSを含む)
3. インタープリター言語の場合、言語のインストール
4. 依存ライブラリのインストール
5. アプリケーションの配置
6. 設定 (環境変数や設定ファイル)
7. 実行
8. サービスに組み込み (LBなど)

#### 十分に仮想化された場合 (Docker)

1. Dockerfile
    - 依存パッケージのインストール (OSを含む)
    - インタープリター言語の場合、言語ランタイムのインストール
    - 依存ライブラリのインストール
7. 実行
    - 設定 (環境変数や設定ファイル)
    - リソースを確保 (VMやDockerで)
8. サービスに組み込み (LBなど)

#### ベアメタルの場合

1. ホストの起動
2. OSのインストール
3. ネットワークの疎通性確保
4. 依存パッケージのインストール (OSを含む)
3. インタープリター言語の場合、言語のインストール
4. 依存ライブラリのインストール
5. アプリケーションの配置
6. 設定 (環境変数や設定ファイル)
7. 実行
8. サービスに組み込み (LBなど)

#### 更新作業
##### ブルーグリーンデプロイ (immutable)

1. 新規のリソースを確保 (VMやDockerで)
2. 依存パッケージのインストール (OSを含む)
3. インタープリター言語の場合、言語のインストール
4. 依存ライブラリのインストール
5. アプリケーションの配置
6. 設定 (環境変数や設定ファイル)
7. 実行
8. サービスに組み込み (LBなど)
9. 古いものをサービスから退役

##### (mutable)

1. サービスから退役
2. 動的なリソースの変更 (ディスク増設など)
2. 依存パッケージの更新 (OSを含む)
3. インタープリター言語の場合、言語の更新
4. 依存ライブラリの更新
5. アプリケーションの更新
6. 設定 (環境変数や設定ファイル)
7. 実行
8. サービスに組み込み (LBなど)

#### 運用
##### 障害発生

1. 情報収集
    - メトリクス回収
    - SSHログイン
2. サービス復旧を第一目標に1次対応 (30分以内)
    - リソースの操作
        - VMの再起動
        - Volumeの拡張
3. 恒久的な2次対応

#### ネットワーク機器

1. 起動
2. IPの初期設定 (ここまで手作業)
3. サービスを提供するための初期設定
4. サービスを提供するための設定

### 作業の抽象化

作業を簡単にすると以下のようになる。

1. リソースの確保
2. プロビジョニング
4. アプリケーションの設定
5. アプリケーションのデプロイ
6. サービスに組み込み

よって考えられるリソースは以下のようになる。今後CPUなどのリソースとリソース志向のリソースが重複しているため、リソース志向のリソースをモデルと呼称する。

1. リソースの確保
    - resource: 仮想化されたリソース
        - cpu
        - memory
        - ip addresses
        - hw addr
        - networkid
4. プロビジョニング
    - provisioning: 初期設定、クラスタレベルの展開
        - network: ネットワーク
        - compute: VM
        - volume: 2次記憶領域
        - image: テンプレートイメージ
        - node: 物理ノード
5. アプリケーションの設定
    - configuration: 環境変数や設定ファイル、パッケージインストールなど
        - cloudinit: ?
        - ansible: ?
    - middleware: DBサービスなど
         - dhcp: computeかVMにタグをつける
         - rdb:
6. アプリケーションのデプロイ
    - application: アプリケーションの実行
        - snapshot: ?
        - pod: ?
    - deployment: スケールアウトなどアプリケーションの使い方によって定義
        - job
        - autoscale
        - simple
7. サービスに組み込み
    - service: LBやAPI ゲートウェイに組み込むなど
    - delivery: ブルーグリーンデプロイやカナリアデプロイなど

## ユーザーと権限

### 考えられるユーザー

- n0stack開発者
- n0stack管理者
- n0stack利用者 (インフラエンジニア)
- n0stack利用者 (VPSとしての利用者)
- n0stack利用者 (アプリケーションエンジニア)