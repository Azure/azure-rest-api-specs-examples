
import com.azure.resourcemanager.oracledatabase.models.GenerateAutonomousDatabaseWalletDetails;
import com.azure.resourcemanager.oracledatabase.models.GenerateType;

/**
 * Samples for AutonomousDatabases GenerateWallet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/autonomousDatabase_generateWallet.json
     */
    /**
     * Sample code: AutonomousDatabases_generateWallet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        autonomousDatabasesGenerateWallet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabases()
            .generateWalletWithResponse(
                "rg000", "databasedb1", new GenerateAutonomousDatabaseWalletDetails()
                    .withGenerateType(GenerateType.SINGLE).withIsRegional(false).withPassword("fakeTokenPlaceholder"),
                com.azure.core.util.Context.NONE);
    }
}
