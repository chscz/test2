# window
  - chocolety install
    - powershell(관리자 권한 실행)에서 아래 command 실행하여 설치(https://chocolatey.org/install)
      ```shell
      Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
      ``` 
  - git install
    ```shell
    choco install git
    ```
  - go install 
    ```shell
    choco install go
    ```
  - mariadb install
    - https://mariadb.org/download/?t=mariadb&p=mariadb&r=11.0.4&os=windows&cpu=x86_64&pkg=msi&m=blendbyte
    ```shell
    choco install mariadb --version=10.2.14
    ```
    - set user
      - root / 1111
    - set schema
      ```
      CREATE DATABASE `ab180` /*!40100 COLLATE 'utf8mb4_general_ci' */;
      
      USE `ab180`;
    
      CREATE TABLE `short_link` (
      `id` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_general_ci',
      `created_at` TIMESTAMP NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
      `url` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_general_ci',
      PRIMARY KEY (`id`) USING BTREE
      )
      COLLATE='utf8mb4_general_ci'
      ENGINE=InnoDB;
      ``` 
  - influxdb install
    - https://docs.influxdata.com/influxdb/v2/install/?t=Windows
    ```shell
    choco install influxdb2
    ```
    - set token, org, bucket, measurement
  - grafana install
    - https://grafana.com/grafana/download?platform=windows
      ```shell
      choco install grafana --version=6.1.6
      ````
      ```
        from(bucket: "ab180")
        |> range(start: -7d)
        |> filter(fn: (r) => r["_measurement"] == "data")
        |> aggregateWindow(every: 1d, fn: count)
        |> yield(name: "_value")
      ```
    
  - project root path 에서
    - go mod init ab180
    - go mod tidy
    - go mod vendor
    - go run main.go

# mac
  - homebrew install
  - go install
  - mariadb install
  - influxdb install
  - grafana install

# todo
  - 매크로방지(세션별 지정시간 내의 일정횟수 이상 요청시?)
  - 
