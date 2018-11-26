using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using UnityEngine;

class Entity
{
    protected int mID = 0;

    protected GameObject mObj = null;

    public int id { get { return mID; } }

    public GameObject go { get { return mObj; } }
}

