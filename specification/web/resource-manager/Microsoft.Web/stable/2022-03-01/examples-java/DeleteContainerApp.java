import com.azure.core.util.Context;

/** Samples for ContainerApps Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/DeleteContainerApp.json
     */
    /**
     * Sample code: Delete Container App.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteContainerApp(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getContainerApps().delete("rg", "testWorkerApp0", Context.NONE);
    }
}
