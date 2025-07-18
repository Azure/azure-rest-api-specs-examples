
import com.azure.resourcemanager.dashboard.models.IntegrationFabricProperties;
import java.util.Arrays;

/**
 * Samples for IntegrationFabrics Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/IntegrationFabrics_Create.json
     */
    /**
     * Sample code: IntegrationFabrics_Create.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void integrationFabricsCreate(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.integrationFabrics().define("sampleIntegration").withRegion("West US")
            .withExistingGrafana("myResourceGroup", "myWorkspace")
            .withProperties(new IntegrationFabricProperties().withTargetResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myAks")
                .withDataSourceResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Monitor/accounts/myAmw")
                .withScenarios(Arrays.asList("scenario1", "scenario2")))
            .create();
    }
}
