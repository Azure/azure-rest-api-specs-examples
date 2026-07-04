
import com.azure.resourcemanager.network.fluent.models.NetworkManagerInner;
import com.azure.resourcemanager.network.models.ConfigurationType;
import com.azure.resourcemanager.network.models.NetworkManagerPropertiesNetworkManagerScopes;
import java.util.Arrays;

/**
 * Samples for NetworkManagers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerPut.json
     */
    /**
     * Sample code: Put Network Manager.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void putNetworkManager(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkManagers().createOrUpdateWithResponse("rg1", "TestNetworkManager",
            new NetworkManagerInner().withDescription("My Test Network Manager")
                .withNetworkManagerScopes(new NetworkManagerPropertiesNetworkManagerScopes()
                    .withManagementGroups(Arrays.asList("/Microsoft.Management/testmg"))
                    .withSubscriptions(Arrays.asList("/subscriptions/00000000-0000-0000-0000-000000000000")))
                .withNetworkManagerScopeAccesses(Arrays.asList(ConfigurationType.CONNECTIVITY)),
            com.azure.core.util.Context.NONE);
    }
}
