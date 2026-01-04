
import com.azure.resourcemanager.appservice.models.VirtualNetworkProfile;

/**
 * Samples for AppServiceEnvironments ChangeVnet.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * AppServiceEnvironments_ChangeVnet.json
     */
    /**
     * Sample code: Move an App Service Environment to a different VNET.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        moveAnAppServiceEnvironmentToADifferentVNET(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().changeVnet("test-rg", "test-ase",
            new VirtualNetworkProfile().withId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
            com.azure.core.util.Context.NONE);
    }
}
