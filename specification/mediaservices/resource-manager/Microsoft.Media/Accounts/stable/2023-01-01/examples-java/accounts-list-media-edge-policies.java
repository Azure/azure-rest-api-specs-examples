import com.azure.resourcemanager.mediaservices.models.ListEdgePoliciesInput;

/** Samples for Mediaservices ListEdgePolicies. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2023-01-01/examples/accounts-list-media-edge-policies.json
     */
    /**
     * Sample code: List the media edge policies.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listTheMediaEdgePolicies(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .mediaservices()
            .listEdgePoliciesWithResponse(
                "contosorg",
                "contososports",
                new ListEdgePoliciesInput().withDeviceId("contosiothubhost_contosoiotdevice"),
                com.azure.core.util.Context.NONE);
    }
}
