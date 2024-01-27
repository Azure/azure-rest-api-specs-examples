
import com.azure.core.util.Context;

/** Samples for SyncGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SyncGroupDelete.json
     */
    /**
     * Sample code: Delete a sync group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteASyncGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSyncGroups().delete("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", "syncgroupcrud-3187", Context.NONE);
    }
}
