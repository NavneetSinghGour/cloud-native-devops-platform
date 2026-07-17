async function loadApplicationInfo() {
    try {
        const response = await fetch("/api/info");

        if (!response.ok) {
            throw new Error("Failed to fetch application info");
        }

        const data = await response.json();

        document.getElementById("app-name").textContent = data.AppName;
        document.getElementById("app-version").textContent = data.Version;
        document.getElementById("app-environment").textContent = data.Environment;
        document.getElementById("app-hostname").textContent = data.Hostname;
        document.getElementById("app-time").textContent = data.CurrentTime;
        document.getElementById("app-uptime").textContent = data.Uptime;

    } catch (error) {
        console.error("Application API:", error);
    }
}

async function loadBuildInfo() {
    try {
        const response = await fetch("/api/build");

        if (!response.ok) {
            throw new Error("Failed to fetch build information");
        }

        const data = await response.json();

        document.getElementById("build-version").textContent = data.version;
        document.getElementById("git-commit").textContent = data.commit;
        document.getElementById("git-branch").textContent = data.branch;
        document.getElementById("build-time").textContent = data.buildTime;

    } catch (error) {
        console.error("Build API:", error);
    }
}

async function loadRuntimeInfo() {
    try {
        const response = await fetch("/api/runtime");

        if (!response.ok) {
            throw new Error("Failed to fetch runtime information");
        }

        const data = await response.json();

        document.getElementById("runtime-go").textContent = data.goVersion;
        document.getElementById("runtime-os").textContent = data.os;
        document.getElementById("runtime-arch").textContent = data.arch;
        document.getElementById("runtime-cpu").textContent = data.cpu;
        document.getElementById("runtime-goroutines").textContent = data.goroutines;
        document.getElementById("runtime-memory").textContent =
            (data.memoryAlloc / 1024 / 1024).toFixed(2) + " MB";

    } catch (error) {
        console.error("Runtime API:", error);
    }
}

async function loadKubernetesInfo() {
    try {
        const response = await fetch("/api/kubernetes");

        if (!response.ok) {
            throw new Error("Failed to fetch Kubernetes information");
        }

        const data = await response.json();

        document.getElementById("k8s-pod").textContent = data.podName;
        document.getElementById("k8s-namespace").textContent = data.namespace;
        document.getElementById("k8s-node").textContent = data.nodeName;
        document.getElementById("k8s-podip").textContent = data.podIP;
        document.getElementById("k8s-hostip").textContent = data.hostIP;

    } catch (error) {
        console.error("Kubernetes API:", error);
    }
}

async function refreshDashboard() {
    await Promise.all([
        loadApplicationInfo(),
        loadBuildInfo(),
        loadRuntimeInfo(),
        loadKubernetesInfo()
    ]);
}

refreshDashboard();
setInterval(refreshDashboard, 5000);
