
import com.azure.resourcemanager.network.models.AvailableProvidersListParameters;
import java.util.Arrays;

/**
 * Samples for NetworkWatchers ListAvailableProviders.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkWatcherAvailableProvidersListGet.json
     */
    /**
     * Sample code: Get Available Providers List.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAvailableProvidersList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkWatchers().listAvailableProviders("rg1", "nw1",
            new AvailableProvidersListParameters().withAzureLocations(Arrays.asList("West US"))
                .withCountry("United States").withState("washington").withCity("seattle"),
            com.azure.core.util.Context.NONE);
    }
}
