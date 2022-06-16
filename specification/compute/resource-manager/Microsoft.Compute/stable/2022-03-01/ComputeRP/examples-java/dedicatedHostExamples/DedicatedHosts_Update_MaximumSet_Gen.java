import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.DedicatedHostLicenseTypes;
import com.azure.resourcemanager.compute.models.DedicatedHostUpdate;
import java.util.HashMap;
import java.util.Map;

/** Samples for DedicatedHosts Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/dedicatedHostExamples/DedicatedHosts_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHosts_Update_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dedicatedHostsUpdateMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDedicatedHosts()
            .update(
                "rgcompute",
                "aaaaaaaaa",
                "aaaaaaaaaaaaaaaaaaaaa",
                new DedicatedHostUpdate()
                    .withTags(mapOf("key8813", "aaaaaaaaaaaaaaaaaaaaaaaaaaa"))
                    .withPlatformFaultDomain(1)
                    .withAutoReplaceOnFailure(true)
                    .withLicenseType(DedicatedHostLicenseTypes.WINDOWS_SERVER_HYBRID),
                Context.NONE);
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
