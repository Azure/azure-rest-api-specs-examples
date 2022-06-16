import com.azure.core.util.Context;

/** Samples for Extensions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-01-01/examples/DeleteExtension.json
     */
    /**
     * Sample code: Delete Arc Extension.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteArcExtension(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.extensions().delete("test-rg", "myCluster", "default", "MicrosoftMonitoringAgent", Context.NONE);
    }
}
