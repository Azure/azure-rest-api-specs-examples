
import com.azure.resourcemanager.workloadssapvirtualinstance.models.ManagedResourcesNetworkAccessType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SAPVirtualInstanceIdentity;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SAPVirtualInstanceIdentityType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapVirtualInstance;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.UpdateSapVirtualInstanceProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SapVirtualInstances Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapVirtualInstances_UpdateTrustedAccess.json
     */
    /**
     * Sample code: SAPVirtualInstances_TrustedAccessEnable_Update.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPVirtualInstancesTrustedAccessEnableUpdate(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        SapVirtualInstance resource = manager.sapVirtualInstances()
            .getByResourceGroupWithResponse("test-rg", "X00", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder"))
            .withIdentity(new SAPVirtualInstanceIdentity().withType(SAPVirtualInstanceIdentityType.NONE))
            .withProperties(new UpdateSapVirtualInstanceProperties()
                .withManagedResourcesNetworkAccessType(ManagedResourcesNetworkAccessType.PRIVATE))
            .apply();
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
