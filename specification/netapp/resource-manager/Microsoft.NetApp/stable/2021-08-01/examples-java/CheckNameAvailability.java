import com.azure.core.util.Context;
import com.azure.resourcemanager.netapp.models.CheckNameResourceTypes;
import com.azure.resourcemanager.netapp.models.ResourceNameAvailabilityRequest;

/** Samples for NetAppResource CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/CheckNameAvailability.json
     */
    /**
     * Sample code: CheckNameAvailability.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void checkNameAvailability(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager
            .netAppResources()
            .checkNameAvailabilityWithResponse(
                "eastus",
                new ResourceNameAvailabilityRequest()
                    .withName("accName")
                    .withType(CheckNameResourceTypes.fromString("netAppAccount"))
                    .withResourceGroup("myRG"),
                Context.NONE);
    }
}
