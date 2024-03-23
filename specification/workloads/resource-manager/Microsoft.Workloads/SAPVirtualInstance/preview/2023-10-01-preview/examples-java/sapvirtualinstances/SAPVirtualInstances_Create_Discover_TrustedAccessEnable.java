
import com.azure.resourcemanager.workloadssapvirtualinstance.models.DiscoveryConfiguration;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.ManagedResourcesNetworkAccessType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapEnvironmentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapProductType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapVirtualInstanceProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SapVirtualInstances Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/
     * examples/sapvirtualinstances/SAPVirtualInstances_Create_Discover_TrustedAccessEnable.json
     */
    /**
     * Sample code: Register with trusted access enabled.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void registerWithTrustedAccessEnabled(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().define("X00").withRegion("northeurope").withExistingResourceGroup("test-rg")
            .withProperties(new SapVirtualInstanceProperties().withEnvironment(SapEnvironmentType.NON_PROD)
                .withSapProduct(SapProductType.S4HANA)
                .withManagedResourcesNetworkAccessType(ManagedResourcesNetworkAccessType.PRIVATE)
                .withConfiguration(new DiscoveryConfiguration().withCentralServerVmId(
                    "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/sapq20scsvm0")))
            .withTags(mapOf("createdby", "abc@microsoft.com", "test", "abc")).create();
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
