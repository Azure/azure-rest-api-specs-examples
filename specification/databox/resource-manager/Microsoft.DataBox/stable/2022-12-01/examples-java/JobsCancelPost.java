import com.azure.resourcemanager.databox.models.CancellationReason;

/** Samples for Jobs Cancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsCancelPost.json
     */
    /**
     * Sample code: JobsCancelPost.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsCancelPost(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager
            .jobs()
            .cancelWithResponse(
                "YourResourceGroupName",
                "TestJobName1",
                new CancellationReason().withReason("CancelTest"),
                com.azure.core.util.Context.NONE);
    }
}
