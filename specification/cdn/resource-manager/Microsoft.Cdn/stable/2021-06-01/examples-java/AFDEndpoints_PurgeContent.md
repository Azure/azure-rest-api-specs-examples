```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.models.AfdPurgeParameters;
import java.util.Arrays;

/** Samples for AfdEndpoints PurgeContent. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDEndpoints_PurgeContent.json
     */
    /**
     * Sample code: AFDEndpoints_PurgeContent.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDEndpointsPurgeContent(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getAfdEndpoints()
            .purgeContent(
                "RG",
                "profile1",
                "endpoint1",
                new AfdPurgeParameters()
                    .withContentPaths(Arrays.asList("/folder1"))
                    .withDomains(Arrays.asList("endpoint1.azureedge.net")),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
