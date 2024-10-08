
import com.azure.resourcemanager.azurestackhci.fluent.models.UpdateSummariesInner;
import com.azure.resourcemanager.azurestackhci.models.UpdateSummariesPropertiesState;
import java.time.OffsetDateTime;

/**
 * Samples for UpdateSummariesOperation Put.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * PutUpdateSummaries.json
     */
    /**
     * Sample code: Put Update summaries under cluster resource.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        putUpdateSummariesUnderClusterResource(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.updateSummariesOperations().putWithResponse("testrg", "testcluster",
            new UpdateSummariesInner().withOemFamily("DellEMC").withHardwareModel("PowerEdge R730xd")
                .withCurrentVersion("4.2203.2.32").withLastUpdated(OffsetDateTime.parse("2022-04-06T14:08:18.254Z"))
                .withLastChecked(OffsetDateTime.parse("2022-04-07T18:04:07Z"))
                .withState(UpdateSummariesPropertiesState.APPLIED_SUCCESSFULLY),
            com.azure.core.util.Context.NONE);
    }
}
