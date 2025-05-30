
import com.azure.resourcemanager.dynatrace.fluent.models.DynatraceSingleSignOnProperties;
import com.azure.resourcemanager.dynatrace.models.AccountInfo;
import com.azure.resourcemanager.dynatrace.models.DynatraceEnvironmentProperties;
import com.azure.resourcemanager.dynatrace.models.EnvironmentInfo;
import com.azure.resourcemanager.dynatrace.models.IdentityProperties;
import com.azure.resourcemanager.dynatrace.models.ManagedIdentityType;
import com.azure.resourcemanager.dynatrace.models.MarketplaceSubscriptionStatus;
import com.azure.resourcemanager.dynatrace.models.MonitoringStatus;
import com.azure.resourcemanager.dynatrace.models.PlanData;
import com.azure.resourcemanager.dynatrace.models.UserInfo;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Monitors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/
     * Monitors_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_CreateOrUpdate_MaximumSet_Gen.
     * 
     * @param manager Entry point to DynatraceManager.
     */
    public static void
        monitorsCreateOrUpdateMaximumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.monitors().define("myMonitor").withRegion("West US 2").withExistingResourceGroup("myResourceGroup")
            .withTags(mapOf("Environment", "Dev"))
            .withIdentity(new IdentityProperties().withType(ManagedIdentityType.SYSTEM_ASSIGNED))
            .withMonitoringStatus(MonitoringStatus.ENABLED)
            .withMarketplaceSubscriptionStatus(MarketplaceSubscriptionStatus.ACTIVE)
            .withDynatraceEnvironmentProperties(new DynatraceEnvironmentProperties().withAccountInfo(new AccountInfo())
                .withEnvironmentInfo(new EnvironmentInfo())
                .withSingleSignOnProperties(new DynatraceSingleSignOnProperties()))
            .withUserInfo(new UserInfo().withFirstName("Alice").withLastName("Bobab")
                .withEmailAddress("alice@microsoft.com").withPhoneNumber("123456").withCountry("westus2"))
            .withPlanData(new PlanData().withUsageType("Committed").withBillingCycle("Monthly")
                .withPlanDetails("dynatraceapitestplan")
                .withEffectiveDate(OffsetDateTime.parse("2019-08-30T15:14:33+02:00")))
            .create();
    }

    // Use "Map.of" if available
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
