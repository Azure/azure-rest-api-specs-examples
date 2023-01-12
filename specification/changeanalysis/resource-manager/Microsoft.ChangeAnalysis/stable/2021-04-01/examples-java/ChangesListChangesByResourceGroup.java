import java.time.OffsetDateTime;

/** Samples for Changes ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/changeanalysis/resource-manager/Microsoft.ChangeAnalysis/stable/2021-04-01/examples/ChangesListChangesByResourceGroup.json
     */
    /**
     * Sample code: Changes_ListChangesByResourceGroup.
     *
     * @param manager Entry point to AzureChangeAnalysisManager.
     */
    public static void changesListChangesByResourceGroup(
        com.azure.resourcemanager.changeanalysis.AzureChangeAnalysisManager manager) {
        manager
            .changes()
            .listByResourceGroup(
                "MyResourceGroup",
                OffsetDateTime.parse("2021-04-25T12:09:03.141Z"),
                OffsetDateTime.parse("2021-04-26T12:09:03.141Z"),
                null,
                com.azure.core.util.Context.NONE);
    }
}
