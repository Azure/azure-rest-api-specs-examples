Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.AvailableProvidersListParameters;
import java.util.Arrays;

/** Samples for NetworkWatchers ListAvailableProviders. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherAvailableProvidersListGet.json
     */
    /**
     * Sample code: Get Available Providers List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAvailableProvidersList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkWatchers()
            .listAvailableProviders(
                "rg1",
                "nw1",
                new AvailableProvidersListParameters()
                    .withAzureLocations(Arrays.asList("West US"))
                    .withCountry("United States")
                    .withState("washington")
                    .withCity("seattle"),
                Context.NONE);
    }
}
```
