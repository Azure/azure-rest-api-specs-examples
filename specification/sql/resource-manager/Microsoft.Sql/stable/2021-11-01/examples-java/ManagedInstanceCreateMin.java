
import com.azure.resourcemanager.sql.fluent.models.ManagedInstanceInner;
import com.azure.resourcemanager.sql.models.ManagedInstanceLicenseType;
import com.azure.resourcemanager.sql.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ManagedInstances CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceCreateMin.json
     */
    /**
     * Sample code: Create managed instance with minimal properties.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createManagedInstanceWithMinimalProperties(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstances().createOrUpdate("testrg", "testinstance",
            new ManagedInstanceInner().withLocation("Japan East")
                .withSku(new Sku().withName("GP_Gen4").withTier("GeneralPurpose")).withAdministratorLogin("dummylogin")
                .withAdministratorLoginPassword("fakeTokenPlaceholder")
                .withSubnetId(
                    "/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1")
                .withLicenseType(ManagedInstanceLicenseType.LICENSE_INCLUDED).withVCores(8).withStorageSizeInGB(1024),
            com.azure.core.util.Context.NONE);
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
