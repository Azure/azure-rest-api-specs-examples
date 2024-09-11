
import com.azure.resourcemanager.qumulo.models.MarketplaceDetails;
import com.azure.resourcemanager.qumulo.models.UserDetails;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for FileSystems CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2024-06-19/examples/
     * FileSystems_CreateOrUpdate_MinimumSet_Gen.json
     */
    /**
     * Sample code: FileSystems_CreateOrUpdate_MinimumSet_Gen.
     * 
     * @param manager Entry point to QumuloManager.
     */
    public static void fileSystemsCreateOrUpdateMinimumSetGen(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        manager.fileSystems().define("aaaaaaaa").withRegion("aaaaaaaaaaaaaaaaaaaaaaaaa")
            .withExistingResourceGroup("rgopenapi")
            .withMarketplaceDetails(new MarketplaceDetails().withMarketplaceSubscriptionId("aaaaaaaaaaaaa")
                .withPlanId("aaaaaa").withOfferId("aaaaaaaaaaaaaaaaaaaaaaaaa"))
            .withStorageSku("Standard").withUserDetails(new UserDetails().withEmail("viptslwulnpaupfljvnjeq"))
            .withDelegatedSubnetId("aaaaaaaaaa").withAdminPassword("fakeTestSecretPlaceholder").create();
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
