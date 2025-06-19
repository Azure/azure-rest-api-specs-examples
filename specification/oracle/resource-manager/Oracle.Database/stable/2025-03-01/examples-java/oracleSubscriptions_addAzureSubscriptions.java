
import com.azure.resourcemanager.oracledatabase.models.AzureSubscriptions;
import java.util.Arrays;

/**
 * Samples for OracleSubscriptions AddAzureSubscriptions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/oracleSubscriptions_addAzureSubscriptions.json
     */
    /**
     * Sample code: OracleSubscriptions_addAzureSubscriptions.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void oracleSubscriptionsAddAzureSubscriptions(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.oracleSubscriptions().addAzureSubscriptions(
            new AzureSubscriptions().withAzureSubscriptionIds(Arrays.asList("00000000-0000-0000-0000-000000000001")),
            com.azure.core.util.Context.NONE);
    }
}
