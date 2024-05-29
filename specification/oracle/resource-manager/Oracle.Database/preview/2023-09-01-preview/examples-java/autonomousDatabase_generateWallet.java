
import com.azure.resourcemanager.oracledatabase.models.GenerateAutonomousDatabaseWalletDetails;
import com.azure.resourcemanager.oracledatabase.models.GenerateType;

/**
 * Samples for AutonomousDatabases GenerateWallet.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/
     * autonomousDatabase_generateWallet.json
     */
    /**
     * Sample code: Generate wallet action on Autonomous Database.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void generateWalletActionOnAutonomousDatabase(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabases()
            .generateWalletWithResponse(
                "rg000", "databasedb1", new GenerateAutonomousDatabaseWalletDetails()
                    .withGenerateType(GenerateType.SINGLE).withIsRegional(false).withPassword("fakeTokenPlaceholder"),
                com.azure.core.util.Context.NONE);
    }
}
