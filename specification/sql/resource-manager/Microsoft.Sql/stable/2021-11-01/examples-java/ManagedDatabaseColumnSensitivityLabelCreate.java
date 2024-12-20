
import com.azure.resourcemanager.sql.fluent.models.SensitivityLabelInner;
import com.azure.resourcemanager.sql.models.SensitivityLabelRank;

/**
 * Samples for ManagedDatabaseSensitivityLabels CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedDatabaseColumnSensitivityLabelCreate.json
     */
    /**
     * Sample code: Updates or creates a sensitivity label of a given column with all parameters in a managed database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updatesOrCreatesASensitivityLabelOfAGivenColumnWithAllParametersInAManagedDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseSensitivityLabels().createOrUpdateWithResponse(
            "myRG", "myManagedInstanceName", "myDatabase", "dbo", "myTable", "myColumn",
            new SensitivityLabelInner().withLabelName("PII").withLabelId("bf91e08c-f4f0-478a-b016-25164b2a65ff")
                .withInformationType("PhoneNumber").withInformationTypeId("d22fa6e9-5ee4-3bde-4c2b-a409604c4646")
                .withRank(SensitivityLabelRank.HIGH),
            com.azure.core.util.Context.NONE);
    }
}
