/** Samples for ApplicationTypeVersions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationTypeVersionDeleteOperation_example.json
     */
    /**
     * Sample code: Delete an application type version.
     *
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void deleteAnApplicationTypeVersion(
        com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager
            .applicationTypeVersions()
            .delete("resRg", "myCluster", "myAppType", "1.0", com.azure.core.util.Context.NONE);
    }
}
