import com.azure.resourcemanager.timeseriesinsights.models.AccessPolicyRole;
import java.util.Arrays;

/** Samples for AccessPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/AccessPoliciesCreate.json
     */
    /**
     * Sample code: AccessPoliciesCreate.
     *
     * @param manager Entry point to TimeSeriesInsightsManager.
     */
    public static void accessPoliciesCreate(
        com.azure.resourcemanager.timeseriesinsights.TimeSeriesInsightsManager manager) {
        manager
            .accessPolicies()
            .define("ap1")
            .withExistingEnvironment("rg1", "env1")
            .withPrincipalObjectId("aGuid")
            .withDescription("some description")
            .withRoles(Arrays.asList(AccessPolicyRole.READER))
            .create();
    }
}
