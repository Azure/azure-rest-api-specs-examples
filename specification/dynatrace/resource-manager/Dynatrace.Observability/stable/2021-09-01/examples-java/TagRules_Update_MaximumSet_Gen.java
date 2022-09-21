import com.azure.core.util.Context;
import com.azure.resourcemanager.dynatrace.models.FilteringTag;
import com.azure.resourcemanager.dynatrace.models.LogRules;
import com.azure.resourcemanager.dynatrace.models.MetricRules;
import com.azure.resourcemanager.dynatrace.models.SendAadLogsStatus;
import com.azure.resourcemanager.dynatrace.models.SendActivityLogsStatus;
import com.azure.resourcemanager.dynatrace.models.SendSubscriptionLogsStatus;
import com.azure.resourcemanager.dynatrace.models.TagAction;
import com.azure.resourcemanager.dynatrace.models.TagRule;
import java.util.Arrays;

/** Samples for TagRules Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/TagRules_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: TagRules_Update_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void tagRulesUpdateMaximumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        TagRule resource =
            manager.tagRules().getWithResponse("myResourceGroup", "myMonitor", "default", Context.NONE).getValue();
        resource
            .update()
            .withLogRules(
                new LogRules()
                    .withSendAadLogs(SendAadLogsStatus.ENABLED)
                    .withSendSubscriptionLogs(SendSubscriptionLogsStatus.ENABLED)
                    .withSendActivityLogs(SendActivityLogsStatus.ENABLED)
                    .withFilteringTags(
                        Arrays
                            .asList(
                                new FilteringTag()
                                    .withName("Environment")
                                    .withValue("Prod")
                                    .withAction(TagAction.INCLUDE),
                                new FilteringTag()
                                    .withName("Environment")
                                    .withValue("Dev")
                                    .withAction(TagAction.EXCLUDE))))
            .withMetricRules(
                new MetricRules()
                    .withFilteringTags(
                        Arrays
                            .asList(
                                new FilteringTag()
                                    .withName("Environment")
                                    .withValue("Prod")
                                    .withAction(TagAction.INCLUDE))))
            .apply();
    }
}
