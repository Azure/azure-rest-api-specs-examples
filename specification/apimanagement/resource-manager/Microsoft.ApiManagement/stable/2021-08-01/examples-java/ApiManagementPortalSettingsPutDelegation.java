import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.fluent.models.PortalDelegationSettingsInner;
import com.azure.resourcemanager.apimanagement.models.RegistrationDelegationSettingsProperties;
import com.azure.resourcemanager.apimanagement.models.SubscriptionsDelegationSettingsProperties;

/** Samples for DelegationSettings CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementPortalSettingsPutDelegation.json
     */
    /**
     * Sample code: ApiManagementPortalSettingsUpdateDelegation.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementPortalSettingsUpdateDelegation(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .delegationSettings()
            .createOrUpdateWithResponse(
                "rg1",
                "apimService1",
                new PortalDelegationSettingsInner()
                    .withUrl("http://contoso.com/delegation")
                    .withValidationKey("<validationKey>")
                    .withSubscriptions(new SubscriptionsDelegationSettingsProperties().withEnabled(true))
                    .withUserRegistration(new RegistrationDelegationSettingsProperties().withEnabled(true)),
                "*",
                Context.NONE);
    }
}
