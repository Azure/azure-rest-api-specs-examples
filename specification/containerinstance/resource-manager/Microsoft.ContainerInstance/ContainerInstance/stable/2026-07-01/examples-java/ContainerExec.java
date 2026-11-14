
import com.azure.resourcemanager.containerinstance.models.ContainerExecRequest;
import com.azure.resourcemanager.containerinstance.models.ContainerExecRequestTerminalSize;

/**
 * Samples for Containers ExecuteCommand.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerExec.json
     */
    /**
     * Sample code: ContainerExec.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void containerExec(com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getContainers()
            .executeCommandWithResponse("demo", "demo1", "container1",
                new ContainerExecRequest().withCommand("/bin/bash")
                    .withTerminalSize(new ContainerExecRequestTerminalSize().withRows(12).withCols(12)),
                com.azure.core.util.Context.NONE);
    }
}
