Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_1.1.0-beta.3/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.ListEdgePoliciesInput;

/** Samples for Mediaservices ListEdgePolicies. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/accounts-list-media-edge-policies.json
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
```
