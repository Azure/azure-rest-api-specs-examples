
import com.azure.resourcemanager.devopsinfrastructure.models.AzureDevOpsOrganizationProfile;
import com.azure.resourcemanager.devopsinfrastructure.models.DevOpsAzureSku;
import com.azure.resourcemanager.devopsinfrastructure.models.Organization;
import com.azure.resourcemanager.devopsinfrastructure.models.PoolImage;
import com.azure.resourcemanager.devopsinfrastructure.models.PoolProperties;
import com.azure.resourcemanager.devopsinfrastructure.models.ProvisioningState;
import com.azure.resourcemanager.devopsinfrastructure.models.StatelessAgentProfile;
import com.azure.resourcemanager.devopsinfrastructure.models.VmssFabricProfile;
import java.util.Arrays;

/**
 * Samples for Pools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-19/CreateOrUpdatePool.json
     */
    /**
     * Sample code: Pools_CreateOrUpdate.
     * 
     * @param manager Entry point to DevOpsInfrastructureManager.
     */
    public static void
        poolsCreateOrUpdate(com.azure.resourcemanager.devopsinfrastructure.DevOpsInfrastructureManager manager) {
        manager.pools().define("pool").withRegion("eastus").withExistingResourceGroup("rg")
            .withProperties(new PoolProperties().withProvisioningState(ProvisioningState.SUCCEEDED)
                .withMaximumConcurrency(10)
                .withOrganizationProfile(new AzureDevOpsOrganizationProfile()
                    .withOrganizations(Arrays.asList(new Organization().withUrl("https://mseng.visualstudio.com"))))
                .withAgentProfile(new StatelessAgentProfile())
                .withFabricProfile(new VmssFabricProfile().withSku(new DevOpsAzureSku().withName("Standard_D4ads_v5"))
                    .withImages(Arrays.asList(new PoolImage()
                        .withResourceId("/MicrosoftWindowsServer/WindowsServer/2019-Datacenter/latest"))))
                .withDevCenterProjectResourceId(
                    "/subscriptions/222e81d0-cf38-4dab-baa5-289bf16baaa4/resourceGroups/rg-1es-devcenter/providers/Microsoft.DevCenter/projects/1ES"))
            .create();
    }
}
