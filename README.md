# 🚀 WCP360: The Next-Generation Linux-Native Control Panel

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Built with Go](https://img.shields.io/badge/Made%20with-Go-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Native Linux](https://img.shields.io/badge/Platform-Linux--Native-E34F26?logo=linux)](https://kernel.org)

**WCP360** is a modern, high-performance, security-first hosting control panel designed to replace legacy platforms like cPanel/WHM. Built with **Go**, it leverages a modular, kernel-native infrastructure layer for operators who demand speed, scalability, and zero legacy bloat.

---

## 🌍 Vision

WCP360 is not just a "panel"—it is a high-precision **infrastructure control plane**. We aim to provide the fastest and safest environment for:
* **High-density** shared hosting.
* **Managed WordPress/Laravel** hosting via FrankenPHP.
* **Cloud-native** multi-tenant deployments.
* **Enterprise** resource governance.

---

## ✨ Key Principles

| Principle | Description |
| :--- | :--- |
| **Performance-First** | Event-driven Go backend, Caddy HTTP/3 gateway, and Redis caching. |
| **Security by Design** | Zero-trust architecture, immutable audit logs, and strict kernel isolation. |
| **Modular** | Everything is a module: Web, SSL, DB, DNS, Email, and Backups. |
| **Linux-Native** | Built on `systemd`, `cgroups v2`, `namespaces`, and `SELinux/AppArmor`. |
| **Cluster Ready** | API-first design, ready for multi-node orchestration (v2+). |

---

## 🧩 Core Features

### ⚙️ Core Platform
* **Resource Governance:** Granular CPU/RAM/IO limits via `cgroups v2`.
* **Immutable Logging:** Cryptographically signed audit trails.
* **Job Engine:** Idempotent task processing via `Asynq` + `Redis`.
* **Zero-Trust:** Secure bootstrap (root → non-root daemon transition).

### 🌐 Web Hosting Stack (v1.0)
* **Caddy Gateway:** Native HTTP/3, QUIC, and Automatic HTTPS (ACME v2).
* **FrankenPHP Runtime:** Worker-based execution—faster than PHP-FPM.
* **Modern Defaults:** Brotli compression and tenant-aware routing.

### 🔐 SSL & Security
* **Dynamic Firewall:** Managed via `nftables`.
* **WAF:** Integrated `Coraza` + `OWASP CRS`.
* **Isolation:** Per-tenant users and secure process namespaces.

### 🗄 Database & Modules
* **Auto-Provisioning:** MariaDB & PostgreSQL with per-tenant privileges.
* **Extensible SDK:** Build custom modules for DNS (PowerDNS), Email (Postfix), or S3 Backups.

---

## 🏗 Architecture Overview



WCP360 maintains a minimal, immutable core that manages the lifecycle of tenants and modules.

* **Brain:** RBAC, Tenant lifecycle, and Job orchestration.
* **Transport:** Caddy acts as the single public entry point.
* **Storage:** PostgreSQL for state; Redis for transient data.
* **Kernel:** Direct interaction with `systemd` slices and `cgroups`.

---

## ⚡ Performance Philosophy

We engineer for maximum efficiency even on small VPS instances:
* **Single-binary** Go backend.
* **No PHP-FPM overhead** (using FrankenPHP workers).
* **Lazy service loading** to save memory.
* **Brotli & HTTP/3** enabled by default.

---

## 🐧 System Requirements

| Requirement | Specification |
| :--- | :--- |
| **OS (Tier 1)** | Ubuntu 22.04/24.04+, Debian 12+ |
| **OS (Tier 2)** | AlmaLinux 9+, Rocky Linux 9+ |
| **Kernel** | 5.10+ |
| **Features** | `systemd`, `cgroups v2` enabled |

---

## 🗺 Roadmap

- [x] **v1.0: Single-Node Core** - Stable tenants, Caddy/FrankenPHP, MariaDB.
- [ ] **v1.1: Observability** - Prometheus metrics, advanced logging dashboard.
- [ ] **v2.0: Multi-Node** - Central orchestrator, remote agents, tenant migration.
- [ ] **v3.0: Ecosystem** - Third-party module marketplace & Distributed storage.

---

## 🤝 Contributing

We welcome Go developers, sysadmins, and security engineers! 
1. Read **[CONTRIBUTING.md](CONTRIBUTING.md)** for workflows.
2. Explore **[ARCHITECTURE.md](ARCHITECTURE.md)** for design deep-dives.
3. Check the **[ROADMAP.md](ROADMAP.md)** for open tasks.

**License:** Distributed under the [MIT License](LICENSE).

---
*Let's reinvent web hosting together.*
