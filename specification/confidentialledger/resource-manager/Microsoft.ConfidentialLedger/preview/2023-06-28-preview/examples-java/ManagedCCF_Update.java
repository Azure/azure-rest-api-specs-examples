
import com.azure.resourcemanager.confidentialledger.models.DeploymentType;
import com.azure.resourcemanager.confidentialledger.models.LanguageRuntime;
import com.azure.resourcemanager.confidentialledger.models.ManagedCcf;
import com.azure.resourcemanager.confidentialledger.models.ManagedCcfProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ManagedCcf Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-06-28-preview/
     * examples/ManagedCCF_Update.json
     */
    /**
     * Sample code: ManagedCCFUpdate.
     * 
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void
        managedCCFUpdate(com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        ManagedCcf resource = manager.managedCcfs().getByResourceGroupWithResponse("DummyResourceGroupName",
            "DummyMccfAppName", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("additionalProps1", "additional properties"))
            .withProperties(new ManagedCcfProperties()
                .withDeploymentType(new DeploymentType().withLanguageRuntime(LanguageRuntime.CPP).withAppSourceUri(
                    "https://myaccount.blob.core.windows.net/storage/mccfsource?sv=2022-02-11%st=2022-03-11")))
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
