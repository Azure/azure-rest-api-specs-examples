import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.ListEdgePoliciesInput;

/** Samples for Mediaservices ListEdgePolicies. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/accounts-list-media-edge-policies.json
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
                "contoso",
                "contososports",
                new ListEdgePoliciesInput().withDeviceId("contosiothubhost_contosoiotdevice"),
                Context.NONE);
    }
}
