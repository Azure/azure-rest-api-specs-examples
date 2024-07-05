
import com.azure.resourcemanager.oracledatabase.models.AutonomousDatabaseProperties;
import com.azure.resourcemanager.oracledatabase.models.ComputeModel;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AutonomousDatabases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabase_create.json
     */
    /**
     * Sample code: Create Autonomous Database.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        createAutonomousDatabase(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabases().define("databasedb1").withRegion("eastus").withExistingResourceGroup("rg000")
            .withTags(mapOf("tagK1", "tagV1"))
            .withProperties(new AutonomousDatabaseProperties().withAdminPassword("fakeTokenPlaceholder")
                .withCharacterSet("AL32UTF8").withComputeCount(2.0F).withComputeModel(ComputeModel.ECPU)
                .withDataStorageSizeInTbs(1).withDbVersion("18.4.0.0").withDisplayName("example_autonomous_databasedb1")
                .withNcharacterSet("AL16UTF16")
                .withSubnetId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1")
                .withVnetId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1"))
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
