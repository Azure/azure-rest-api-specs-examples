import java.time.OffsetDateTime;

/** Samples for Changes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/changeanalysis/resource-manager/Microsoft.ChangeAnalysis/stable/2021-04-01/examples/ChangesListChangesBySubscription.json
     */
    /**
     * Sample code: Changes_ListChangesBySubscription.
     *
     * @param manager Entry point to AzureChangeAnalysisManager.
     */
    public static void changesListChangesBySubscription(
        com.azure.resourcemanager.changeanalysis.AzureChangeAnalysisManager manager) {
        manager
            .changes()
            .list(
                OffsetDateTime.parse("2021-04-25T12:09:03.141Z"),
                OffsetDateTime.parse("2021-04-26T12:09:03.141Z"),
                null,
                com.azure.core.util.Context.NONE);
    }
}
