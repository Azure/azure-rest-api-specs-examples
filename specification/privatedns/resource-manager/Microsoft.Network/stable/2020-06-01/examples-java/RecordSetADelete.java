import com.azure.core.util.Context;
import com.azure.resourcemanager.privatedns.models.RecordType;

/** Samples for RecordSets Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetADelete.json
     */
    /**
     * Sample code: DELETE Private DNS Zone A Record Set.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dELETEPrivateDNSZoneARecordSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .privateDnsZones()
            .manager()
            .serviceClient()
            .getRecordSets()
            .deleteWithResponse("resourceGroup1", "privatezone1.com", RecordType.A, "recordA", null, Context.NONE);
    }
}
