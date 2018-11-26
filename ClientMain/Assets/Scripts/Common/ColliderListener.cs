using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System;
using ILRuntime.Runtime.Enviorment;

[RequireComponent(typeof(Collider))]
public class ColliderListener : MonoBehaviour
{
    public Action<Collider> onTriggerEnter = delegate { };
    public Action<Collider> onTriggerStay = delegate { };
    public Action<Collider> onTriggerExit = delegate { };
    public Action<Collision> onCollisionEnter = delegate { };
    public Action<Collision> onCollisionStay = delegate { };
    public Action<Collision> onCollisionExit = delegate { };

    public static void RegisterILRuntime(ILRuntime.Runtime.Enviorment.AppDomain domain)
    {
        domain.DelegateManager.RegisterDelegateConvertor<UnityEngine.Events.UnityAction>((act) =>
        {
            return new UnityEngine.Events.UnityAction(() =>
            {
                ((System.Action)act)();
            });
        });

        domain.DelegateManager.RegisterMethodDelegate<UnityEngine.Collider>();
    }

    private void OnTriggerEnter(Collider other)
    {
        onTriggerEnter(other);
    }

    private void OnTriggerStay(Collider other)
    {
        onTriggerStay(other);
    }

    private void OnTriggerExit(Collider other)
    {
        onTriggerExit(other);
    }

    private void OnCollisionEnter(Collision other)
    {
        onCollisionEnter(other);
    }

    private void OnCollisionStay(Collision other)
    {
        onCollisionStay(other);
    }

    private void OnCollisionExit(Collision other)
    {
        onCollisionExit(other);
    }

    public virtual void Clear()
    {
        onTriggerEnter = delegate { };
        onTriggerStay = delegate { };
        onTriggerExit = delegate { };
        onCollisionEnter = delegate { };
        onCollisionStay = delegate { };
        onCollisionExit = delegate { };
    }
}

