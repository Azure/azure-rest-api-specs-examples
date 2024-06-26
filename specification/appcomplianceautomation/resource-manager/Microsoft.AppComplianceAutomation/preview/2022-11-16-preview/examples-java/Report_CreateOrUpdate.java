import com.azure.core.util.Context;
import com.azure.resourcemanager.appcomplianceautomation.fluent.models.ReportResourceInner;
import com.azure.resourcemanager.appcomplianceautomation.models.ReportProperties;
import com.azure.resourcemanager.appcomplianceautomation.models.ResourceMetadata;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ReportOperation CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Report_CreateOrUpdate.json
     */
    /**
     * Sample code: Report_CreateOrUpdate.
     *
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void reportCreateOrUpdate(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager
            .reportOperations()
            .createOrUpdate(
                "testReportName",
                new ReportResourceInner()
                    .withProperties(
                        new ReportProperties()
                            .withOfferGuid("0000")
                            .withTimeZone("GMT Standard Time")
                            .withTriggerTime(OffsetDateTime.parse("2022-03-04T05:11:56.197Z"))
                            .withResources(
                                Arrays
                                    .asList(
                                        new ResourceMetadata()
                                            .withResourceId(
                                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint")
                                            .withTags(mapOf("key1", "value1"))))),
                Context.NONE);
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
