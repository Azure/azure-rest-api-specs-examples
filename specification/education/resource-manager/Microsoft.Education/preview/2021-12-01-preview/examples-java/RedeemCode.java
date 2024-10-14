
import com.azure.resourcemanager.education.models.RedeemRequest;

/**
 * Samples for ResourceProvider RedeemInvitationCode.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/RedeemCode.json
     */
    /**
     * Sample code: RedeemCode.
     * 
     * @param manager Entry point to EducationManager.
     */
    public static void redeemCode(com.azure.resourcemanager.education.EducationManager manager) {
        manager.resourceProviders().redeemInvitationCodeWithResponse(
            new RedeemRequest().withRedeemCode("fakeTokenPlaceholder").withFirstName("test").withLastName("user"),
            com.azure.core.util.Context.NONE);
    }
}
