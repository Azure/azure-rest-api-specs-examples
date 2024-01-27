
import com.azure.core.util.Context;

/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/
     * PrivateLinkResourcesGet.json
     */
    /**
     * Sample code: NameSpacePrivateLinkResourcesGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpacePrivateLinkResourcesGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getPrivateLinkResources().getWithResponse("ArunMonocle",
            "sdk-Namespace-2924", Context.NONE);
    }
}
