# CiWallet Mono Repo

This is a Mono Repo for the CiWallet project, which consists of multiple modules designed to provide a comprehensive multi-chain wallet experience along with related gaming and server functionalities.

### Important Note
We still update code in other repositories; however, for the Sigma Hackathon, we will not create new commits in this repository. Feel free to check and judge the existing code.
Details README in each project folder.

## Modules

### CiWallet
- **Folder Name:** ciwallet
- **Technologies:** TypeScript, Next.js
- **Description:** The first multi-chain wallet launched on Telegram. With the Wormhole TS SDK, CiWallet fully supports Wormhole features such as transfer, redeem, and more.

---

### CiFarm Game Client
- **Folder Name:** cifarm-client
- **Technologies:** Unity, C#
- **Description:** The CiFarm Client is a Unity-based game application designed to deliver an immersive and interactive gaming experience. Built with C#, this client serves as the primary interface for players.

---

### CiFarm Game Server
- **Folder Name:** cifarm-server
- **Technologies:** Go, Nakama
- **Description:** This server acts as the backbone of the CiFarm game, ensuring that gameplay is engaging, dynamic, and responsive to player actions. It’s built to scale efficiently, supporting multiple concurrent users while maintaining performance and reliability.

---

### CiFarm Periphery Server
- **Folder Name:** cifarm-periphery
- **Technologies:** TypeScript, Nest.js
- **Description:** This peripheral server provides an API that enables CiFarm to connect with every blockchain.

---

### CiWallet Bots
- **Folder Name:** ciwallet-bots
- **Technologies:** Python
- **Description:** The CiWallet Bot is a Telegram bot that allows users to create and manage their wallets directly within Telegram.

Overview Diagram:
![CiWallet Overview](https://violet-lazy-yak-333.mypinata.cloud/ipfs/QmTBYxH6Q4KPHFX97N5Jya54tk93x3onRSL8spBTtWVQS2)
