
import com.azure.resourcemanager.appservice.fluent.models.AppServiceEnvironmentResourceInner;
import com.azure.resourcemanager.appservice.models.VirtualNetworkProfile;

/**
 * Samples for AppServiceEnvironments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        createOrUpdateAnAppServiceEnvironment(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().createOrUpdate("test-rg", "test-ase",
            new AppServiceEnvironmentResourceInner().withLocation("South Central US").withKind("Asev3")
                .withVirtualNetwork(new VirtualNetworkProfile().withId(
                    "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/delegated")),
            com.azure.core.util.Context.NONE);
    }
}
