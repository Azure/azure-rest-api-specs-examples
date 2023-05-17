import com.azure.resourcemanager.appcontainers.fluent.models.JobExecutionBaseInner;
import com.azure.resourcemanager.appcontainers.models.JobExecutionNamesCollection;
import java.util.Arrays;

/** Samples for Jobs StopMultipleExecutions. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/Job_Stop_Multiple.json
     */
    /**
     * Sample code: Terminate Multiple Container Apps Job.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void terminateMultipleContainerAppsJob(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .jobs()
            .stopMultipleExecutions(
                "rg",
                "testcontainerAppsJob0",
                new JobExecutionNamesCollection()
                    .withValue(
                        Arrays
                            .asList(
                                new JobExecutionBaseInner().withName("jobExecution-27944453"),
                                new JobExecutionBaseInner().withName("jobExecution-27944452"),
                                new JobExecutionBaseInner().withName("jobExecution-27944451"))),
                com.azure.core.util.Context.NONE);
    }
}
