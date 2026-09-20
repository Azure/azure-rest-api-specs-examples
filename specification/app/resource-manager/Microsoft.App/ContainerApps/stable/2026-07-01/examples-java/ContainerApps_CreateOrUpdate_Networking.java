
import com.azure.resourcemanager.appcontainers.models.Container;
import com.azure.resourcemanager.appcontainers.models.ContainerAppNetworkingConfiguration;
import com.azure.resourcemanager.appcontainers.models.ContainerResources;
import com.azure.resourcemanager.appcontainers.models.Template;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ContainerApps CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerApps_CreateOrUpdate_Networking.json
     */
    /**
     * Sample code: Create or Update Container App with per-app outbound VNet subnet.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateContainerAppWithPerAppOutboundVNetSubnet(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().define("testcontainerApp0").withRegion("East US").withExistingResourceGroup("rg")
            .withEnvironmentId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/expressenv")
            .withNetworking(new ContainerAppNetworkingConfiguration().withOutboundVnetSubnetId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/appsubnet"))
            .withTemplate(new Template().withContainers(
                Arrays.asList(new Container().withImage("repo/testcontainerApp0:v1").withName("testcontainerApp0")
                    .withResources(new ContainerResources().withCpu(0.2D).withMemory("100Mi")))))
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
