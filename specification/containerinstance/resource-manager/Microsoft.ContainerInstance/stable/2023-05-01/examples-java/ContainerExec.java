import com.azure.resourcemanager.containerinstance.models.ContainerExecRequest;
import com.azure.resourcemanager.containerinstance.models.ContainerExecRequestTerminalSize;

/** Samples for Containers ExecuteCommand. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/ContainerExec.json
     */
    /**
     * Sample code: ContainerExec.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerExec(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerGroups()
            .manager()
            .serviceClient()
            .getContainers()
            .executeCommandWithResponse(
                "demo",
                "demo1",
                "container1",
                new ContainerExecRequest()
                    .withCommand("/bin/bash")
                    .withTerminalSize(new ContainerExecRequestTerminalSize().withRows(12).withCols(12)),
                com.azure.core.util.Context.NONE);
    }
}
