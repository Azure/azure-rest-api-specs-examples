
import com.azure.resourcemanager.netapp.models.CheckNameResourceTypes;
import com.azure.resourcemanager.netapp.models.ResourceNameAvailabilityRequest;

/**
 * Samples for NetAppResource CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/CheckNameAvailability.json
     */
    /**
     * Sample code: CheckNameAvailability.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void checkNameAvailability(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.netAppResources().checkNameAvailabilityWithResponse("eastus",
            new ResourceNameAvailabilityRequest().withName("accName")
                .withType(CheckNameResourceTypes.MICROSOFT_NET_APP_NET_APP_ACCOUNTS).withResourceGroup("myRG"),
            com.azure.core.util.Context.NONE);
    }
}
