import com.azure.core.util.Context;

/** Samples for Global GetDeletedWebAppSnapshots. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetDeletedWebAppSnapshots.json
     */
    /**
     * Sample code: Get Deleted Web App Snapshots.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDeletedWebAppSnapshots(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getGlobals().getDeletedWebAppSnapshotsWithResponse("9", Context.NONE);
    }
}
