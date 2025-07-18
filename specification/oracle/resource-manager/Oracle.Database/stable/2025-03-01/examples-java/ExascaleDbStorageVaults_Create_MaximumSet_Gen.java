
import com.azure.resourcemanager.oracledatabase.models.ExascaleDbStorageInputDetails;
import com.azure.resourcemanager.oracledatabase.models.ExascaleDbStorageVaultProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ExascaleDbStorageVaults Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/ExascaleDbStorageVaults_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExascaleDbStorageVaults_Create_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void exascaleDbStorageVaultsCreateMaximumSet(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.exascaleDbStorageVaults().define("vmClusterName").withRegion("ltguhzffucaytqg")
            .withExistingResourceGroup("rgopenapi").withTags(mapOf("key7827", "fakeTokenPlaceholder"))
            .withProperties(new ExascaleDbStorageVaultProperties().withAdditionalFlashCacheInPercent(0)
                .withDescription("dmnvnnduldfmrmkkvvsdtuvmsmruxzzpsfdydgytlckutfozephjygjetrauvbdfcwmti")
                .withDisplayName(
                    "hbsybtelyvhpalemszcvartlhwvskrnpiveqfblvkdihoytqaotdgsgauvgivzaftfgeiwlyeqzssicwrrnlxtsmeakbcsxabjlt")
                .withHighCapacityDatabaseStorageInput(new ExascaleDbStorageInputDetails().withTotalSizeInGbs(21))
                .withTimeZone(
                    "ltrbozwxjunncicrtzjrpqnqrcjgghohztrdlbfjrbkpenopyldwolslwgrgumjfkyovvkzcuxjujuxtjjzubvqvnhrswnbdgcbslopeofmtepbrrlymqwwszvsglmyuvlcuejshtpokirwklnwpcykhyinjmlqvxtyixlthtdishhmtipbygsayvgqzfrprgppylydlcskbmvwctxifdltippfvsxiughqbojqpqrekxsotnqsk"))
            .withZones(Arrays.asList("qk")).create();
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
