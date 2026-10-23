
import com.azure.resourcemanager.horizondb.fluent.models.MicrosoftEntraAuthenticationInner;
import com.azure.resourcemanager.horizondb.models.AuthenticationState;
import com.azure.resourcemanager.horizondb.models.MicrosoftEntraAuthenticationProperties;

/**
 * Samples for MicrosoftEntraAuthentications CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/MicrosoftEntraAuthentications_CreateOrReplaceEnabled.json
     */
    /**
     * Sample code: Enable Microsoft Entra authentication declaratively.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        enableMicrosoftEntraAuthenticationDeclaratively(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.microsoftEntraAuthentications().createOrReplace("example-resource-group", "example-cluster",
            new MicrosoftEntraAuthenticationInner()
                .withProperties(new MicrosoftEntraAuthenticationProperties().withState(AuthenticationState.ENABLED)),
            com.azure.core.util.Context.NONE);
    }
}
