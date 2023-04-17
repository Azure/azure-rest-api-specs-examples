import com.azure.resourcemanager.workloads.models.DiscoveryConfiguration;
import com.azure.resourcemanager.workloads.models.SapEnvironmentType;
import com.azure.resourcemanager.workloads.models.SapProductType;
import java.util.HashMap;
import java.util.Map;

/** Samples for SapVirtualInstances Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPVirtualInstances_Create_Discover.json
     */
    /**
     * Sample code: Register existing SAP system as Virtual Instance for SAP solutions.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void registerExistingSAPSystemAsVirtualInstanceForSAPSolutions(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .sapVirtualInstances()
            .define("X00")
            .withRegion("northeurope")
            .withExistingResourceGroup("test-rg")
            .withEnvironment(SapEnvironmentType.NON_PROD)
            .withSapProduct(SapProductType.S4HANA)
            .withConfiguration(
                new DiscoveryConfiguration()
                    .withCentralServerVmId(
                        "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/sapq20scsvm0"))
            .withTags(mapOf("createdby", "abc@microsoft.com", "test", "abc"))
            .create();
    }

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
